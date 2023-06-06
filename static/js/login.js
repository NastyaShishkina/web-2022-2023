const inputPassword = document.querySelector(".form__password-box");
const imgEye = document.querySelector(".password-input_show");

imgEye.addEventListener(
  "click",
  () => {
    if (inputPassword.type === "password") {
      inputPassword.type = "text";
      imgEye.src = "/static/img/eye-off.png";
    } else {
      inputPassword.type = "password";
      imgEye.src = "/static/img/eye.png";
    }

  }
)
