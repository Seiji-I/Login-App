import React from "react";

export default class LoginForm extends React.Component{
  render() {
    return(
      <form>
        <div>
          <label>
            USER ID
            <input type="text" value=""/>
          </label>
        </div>
        <div>
          <label>
            PASSWORD
            <input type="password" value=""/>
          </label>
        </div>
        <div>
          <button>Login</button>
        </div>
      </form>
    )
  }
}