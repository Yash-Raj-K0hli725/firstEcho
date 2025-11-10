package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Welcome(c echo.Context) error {
	mHTML := `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Welcome - Online Clinic</title>
</head>
<body style="margin:0; padding:0; height:100vh; background:linear-gradient(135deg, #e0f7fa, #b2dfdb); font-family:'Segoe UI', Arial, sans-serif; display:flex; flex-direction:column; align-items:center; justify-content:center; text-align:center; overflow:hidden;">

  <h1 style="color:#00695c; font-size:42px; margin:0; animation:fadeIn 2s ease-in-out;">
    👩‍⚕️ Welcome to Our Online Clinic
  </h1>

  <p style="color:#004d40; font-size:18px; max-width:500px; margin:20px auto 30px auto; line-height:1.5; animation:slideUp 2.5s ease;">
    Your health, our priority — connect with trusted doctors from the comfort of your home. We’re here 24/7 to take care of you.
  </p>

  <button style="background-color:#009688; color:white; border:none; padding:14px 30px; font-size:16px; border-radius:30px; cursor:pointer; box-shadow:0 4px 10px rgba(0,0,0,0.2); transition:all 0.3s ease; animation:fadeIn 3s ease-in-out;"
    onmouseover="this.style.transform='scale(1.1)'; this.style.backgroundColor='#00796b';"
    onmouseout="this.style.transform='scale(1)'; this.style.backgroundColor='#009688';">
    💚 Get Started
  </button>

  <div style="position:absolute; bottom:20px; color:#004d40; font-size:14px; opacity:0.8;">
    © 2025 Online Clinic
  </div>

  <!-- Inline animations -->
  <style>
    @keyframes fadeIn {
      from { opacity: 0; transform: scale(0.9); }
      to { opacity: 1; transform: scale(1); }
    }
    @keyframes slideUp {
      from { opacity: 0; transform: translateY(40px); }
      to { opacity: 1; transform: translateY(0); }
    }
  </style>

</body>
</html>

`
	return c.HTML(http.StatusOK, mHTML)
}
