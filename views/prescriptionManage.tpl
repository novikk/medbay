<html>
<head>
    <title>Prescriptions</title>
    <link rel="stylesheet" href="/static/css/prescriptions.css">
    <meta name="viewport" content="width=750" />
</head>
<body>
    <div class="header">
    <div class="header_backpresc">Prescriptions</div>
    <div class="header_arrow">
        <img src="/static/img/tracking/arrow.png">
    </div>
    <div class="scan_bg">
        <input type="file" style="width: 0.1px;height: 0.1px;opacity: 0;overflow: hidden;position: absolute;z-index: -1;" id="file" onchange="handleFiles(this.files)"/>
        <label for="file" class="scan_text">SCAN QR</label>
    </div>
    </div>
    <div class="prescription_data">PRESCRIPTION DATA</div>

<div class="form">
    <div class="form_medicine">
        <div class="form_medicine_text">Medicine name</div>
    </div>
    <div class="form_dose">
        <div class="form_dose_text">Dose</div>
    </div>
    <div class="form_cooldown">
        <div class="form_cooldown_text">Cooldown</div>
    </div>
    <div class="sign_bg">
        <div class="sign_text">SIGN</div>
    </div>
    
    <div class="signaturit">
        <img src="/static/img/tracking/signaturit.png">
    </div>

</div>

 <div class="navigation">
    <div class="navigation_tracking">
            <a href="/app/tracking"><img src="/static/img/tracking/tracking_inactive.png"></a>
    </div>
    <div class="navigation_prescriptions">
            <a href="/app/prescription"><img src="/static/img/tracking/prescription_active.png"></a>
    </div>
    <div class="navigation_insights">
            <img src="/static/img/tracking/insights.png">
    </div>
    <div class="navigation_settings">
            <img src="/static/img/tracking/settings.png">
    </div>

    </div>

<script type="text/javascript" src="/static/js/webqr/grid.js"></script>
<script type="text/javascript" src="/static/js/webqr/version.js"></script>
<script type="text/javascript" src="/static/js/webqr/detector.js"></script>
<script type="text/javascript" src="/static/js/webqr/formatinf.js"></script>
<script type="text/javascript" src="/static/js/webqr/errorlevel.js"></script>
<script type="text/javascript" src="/static/js/webqr/bitmat.js"></script>
<script type="text/javascript" src="/static/js/webqr/datablock.js"></script>
<script type="text/javascript" src="/static/js/webqr/bmparser.js"></script>
<script type="text/javascript" src="/static/js/webqr/datamask.js"></script>
<script type="text/javascript" src="/static/js/webqr/rsdecoder.js"></script>
<script type="text/javascript" src="/static/js/webqr/gf256poly.js"></script>
<script type="text/javascript" src="/static/js/webqr/gf256.js"></script>
<script type="text/javascript" src="/static/js/webqr/decoder.js"></script>
<script type="text/javascript" src="/static/js/webqr/qrcode.js"></script>
<script type="text/javascript" src="/static/js/webqr/findpat.js"></script>
<script type="text/javascript" src="/static/js/webqr/alignpat.js"></script>
<script type="text/javascript" src="/static/js/webqr/databr.js"></script>

<script type="text/javascript">
qrcode.callback = read;

function handleFiles(f)
{
	var o=[];
	for(var i =0;i<f.length;i++)
	{
	  var reader = new FileReader();

      reader.onload = (function(theFile) {
        return function(e) {
          qrcode.decode(e.target.result);
        };
      })(f[i]);

      // Read in the image file as a data URL.
      reader.readAsDataURL(f[i]);	}
}
	
function read(a)
{
	alert(a);
}	
</script>
</body>
</html>