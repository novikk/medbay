<html>
<head>
<title>QRCODE</title>

<style type="text/css">
</style>


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
	console.log("Hey")
	alert(a);
}	
</script>

</head>

<body>
<input type="file" onchange="handleFiles(this.files)"/>
</body>

</html>