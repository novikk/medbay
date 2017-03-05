package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/astaxie/beego"
)

type AppController struct {
	beego.Controller
}

type Medicine struct {
	Type       string
	Dose       int
	NumDoses   []int
	Prescribed bool
	Cooldown   int
	LastTake   time.Time
}

type Prescription struct {
	Doctor   string
	Medicine string
	Dose     int
	Cooldown int
}

var addedMedicines []Medicine
var prescriptions []Prescription

var signaturitUrl = ""

func init() {
	// example med
	addedMedicines = append(addedMedicines, Medicine{
		Type:       "Amoxicillin",
		Dose:       500,
		NumDoses:   []int{0, 0, 0},
		Prescribed: true,
		Cooldown:   8,
		LastTake:   time.Now().Add(-60 * 60 * 24 * time.Hour),
	})
}

func (c *AppController) Tracking() {
	c.Data["medicines"] = addedMedicines

	c.TplName = "tracking.tpl"
}

func (c *AppController) Prescription() {
	c.Data["prescriptions"] = prescriptions
	c.Data["signaturit"] = signaturitUrl

	signaturitUrl = ""
	c.TplName = "prescription.tpl"
}

func (c *AppController) PrescriptionManage() {
	c.TplName = "prescriptionManage.tpl"
}

func (c *AppController) TrackingDetail() {
	c.TplName = "trackingDetail.tpl"
}

type SignaturitResp struct {
	CreatedAt string        `json:"created_at"`
	Data      []interface{} `json:"data"`
	ID        string        `json:"id"`
	Documents []struct {
		CreatedAt string        `json:"created_at"`
		Events    []interface{} `json:"events"`
		File      struct {
			Name string `json:"name"`
			Size int    `json:"size"`
		} `json:"file"`
		ID     string `json:"id"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		Status string `json:"status"`
	} `json:"documents"`
	URL string `json:"url"`
}

func (c *AppController) AddPrescription() {
	dose, _ := c.GetInt("dose")
	cd, _ := c.GetInt("cd")

	prescriptions = append(prescriptions, Prescription{
		Doctor:   "Dra. Ramirez",
		Medicine: c.GetString("med"),
		Dose:     dose,
		Cooldown: cd,
	})

	// ASDASD
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	f, err := os.Open("receta.pdf")
	if err != nil {
		println(err)
		return
	}
	defer f.Close()
	fw, err := w.CreateFormFile("files[0]", "receta.pdf")
	if err != nil {
		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		println(err)
		return
	}
	// Add the other fields
	if fw, err = w.CreateFormField("recipients[0][name]"); err != nil {
		println(err)
		return
	}
	if _, err = fw.Write([]byte("John")); err != nil {
		println(err)
		return
	}
	if fw, err = w.CreateFormField("recipients[0][email]"); err != nil {
		println(err)
		return
	}
	if _, err = fw.Write([]byte("novikk7@gmail.com")); err != nil {
		println(err)
		return
	}
	if fw, err = w.CreateFormField("delivery_type"); err != nil {
		println(err)
		return
	}
	if _, err = fw.Write([]byte("url")); err != nil {
		println(err)
		return
	}

	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", "https://api.sandbox.signaturit.com/v3/signatures.json", &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Authorization", "Bearer EdBhopmhsvxQUIUjqwClrGSWfXmgrlSdsJTdrlaumwjxrHwyLoNKGMUqvBUddymChArxXtiqbMRwJVrvUPsEwf")

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		println(err)
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))

	var sr SignaturitResp
	err = json.Unmarshal(body, &sr)

	signaturitUrl = sr.URL

	c.Redirect("/app/prescription", 302)
}
