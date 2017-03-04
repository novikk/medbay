<html>
<head>
    <title>Index</title>
    <link rel="stylesheet" href="/static/css/tracking.css">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1" />
</head>
<body>
    <div class="header">
        <div class="header_title">Tracking</div>
        <div class="header_today">Today</div>
    </div>
    <div class="name">
        <img src="/static/img/tracking/usericon.png" class="name_icon">
        <div class="name_text">Humberto Díaz</div>
    </div>
    <div class="week">
        <img src="/static/img/tracking/week50.png" class="week_image">
    </div>
    <div class="daily">
        <img src="/static/img/tracking/daily50.png" class="daily_50">
    </div>

<div class="prescriptions">
    {{ range .medicines }}
    <div class="prescription">
        <div class="presc_type">{{.Type}} <span class="presc_dose">{{.Dose}}mg</span></div>
        <div class="presc_circles">
            {{ range .NumDoses }}
            <div class="presc_circle"></div>
            {{ end }}
        </div>
    </div>
    {{ end }}
</div>

 <div class="navigation">
    <div class="navigation_tracking">
            <img src="/static/img/tracking/tracking.png">
    </div>
    <div class="navigation_prescriptions">
            <img src="/static/img/tracking/prescriptions.png">
    </div>
    <div class="navigation_insights">
            <img src="/static/img/tracking/insights.png">
    </div>
    <div class="navigation_settings">
            <img src="/static/img/tracking/settings.png">
    </div>

    </div>
</body>
</html>