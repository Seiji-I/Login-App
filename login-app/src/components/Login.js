import React from "react";

export default class LoginForm extends React.Component{
  render() {
    let path = "/login"
    let returnUrl = window.location.protocol + "//" + "localhost:3333"+ path
    console.log(returnUrl)
    return(
      <form action={returnUrl} method="POST" accept-charset="UTF-8">
        <div>
          <label>
            USER NAME
            <input type="text" name="username" />
          </label>
        </div>
        <div>
          <label>
            PASSWORD
            <input type="password" name="password" />
          </label>
        </div>
        <div>
          <button>Login</button>
        </div>
      </form>
    )
  }
}