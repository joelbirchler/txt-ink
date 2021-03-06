package main

import "html/template"

var homeTemplate *template.Template

func init() {
	homeTemplate = template.Must(template.New("home").Parse(`<!doctype html>
  <html lang="en">

  <head>
    <meta charset="utf-8">
    <title>txt-ink</title>
    <style>
      @import url('https://fonts.googleapis.com/css?family=Quattrocento:700&display=swap');

      body, textarea, button {
        font-size: 200%;
        font-family: 'Quattrocento', serif;
        font-weight: bold;
        line-height: 1.5em;
      }

      body {
        margin: 0;
        padding: 1em;
        background-color: white;
      }

      textarea {
        width: 100%;
        font-size: 1em;
        resize: none;
        border: none;
        display: block;
      }

      button {
        display: block;
        background: none;
        box-shadow: none;
        border: 2px solid #000;
        border-radius: 3px;
        font-size: 100%;
        width: 100%;
        background-color: #cdeaff;
      }
    </style>
  </head>

  <body>
    <form id="messageForm" action="/message" method="post">
      <textarea id="textArea" name="message" autofocus="true" rows="3" placeholder="{{.Message}}"></textarea>
      <button id="sendButton" type="submit">send</button>
    </form>
    <script type="text/javascript">
      messageForm = document.getElementById("messageForm")
      textArea = document.getElementById("textArea")

      textArea.addEventListener("keypress", function(e) {
        if (event.which === 13) {
          event.preventDefault();
          messageForm.dispatchEvent(new Event("submit", {cancelable: true}));
        }
      });
    </script>
  </body>
  </html>`))
}
