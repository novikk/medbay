<html>
<head>
    <title>Prescriptionslist</title>
    <link rel="stylesheet" href="/static/css/prescriptionslist.css">
    <meta name="viewport" content="width=750" />
</head>
<body>
    <script type="text/javascript" src="https://code.jquery.com/jquery-3.1.1.min.js"></script>

    <div class="header">
        <div class="header_text">Prescriptions</div>
    <div class="add_bg">
        <a href="/app/prescription/manage"><div class="add_text">ADD</div></a>
    </div>
    </div>

    <div class="list_header">PRESCRIPTIONS</div>
    <div class="list_container">
        {{ range .prescriptions }}
        <div class="presc_list_item">
            <div class="presc_doctor">
                {{ .Doctor }}
            </div>
            <div class="presc_med">
                {{ .Medicine }}
            </div>
        </div>
        {{ end }}
        <div class="presc_list_item">
            <div class="presc_doctor">
                Dr. Smith
            </div>
            <div class="presc_med">
                Ibuprofen
            </div>
        </div>
        <div class="presc_list_item">
            <div class="presc_doctor">
                Dr. Smith
            </div>
            <div class="presc_med">
                Amoxicillin
            </div>
        </div>
        <div class="presc_list_item">
            <div class="presc_doctor">
                Dr. Dickinson
            </div>
            <div class="presc_med">
                Nolotil
            </div>
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
{{ if ne .signaturit "" }}
<script type="text/javascript">$('body').append('<iframe src="{{.signaturit}}" style="position: absolute; top: 0; left: 0; height: 1400px; width: 750px; background: #fff" />')</script>
<!--<script type="text/javascript">
    $(window).on('message', function() {
        $('iframe').remove();
    })
</script>-->
{{ end }}
</body>
</html>