import React from "react";
import axios from "axios";

export default class Login extends React.Component{
  constructor(props) {
    super(props);
    this.state = {
      username: "",
      password: "",
    }
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleChangeName = this.handleChangeName.bind(this);
    this.handleChangePassword = this.handleChangePassword.bind(this);
  }
  render() {
    return(
      <form onSubmit={this.handleSubmit}>
        <div>
          <label>USER NAME
            <input type="text" name="username" onChange={this.handleChangeName}/>
          </label>
       </div>
        <div>
          <label>PASSWORD
            <input type="password" name="password" onChange={this.handleChangePassword}/>
          </label>
        </div>
        <div>
          <button>Login</button>
        </div>
      </form>
    )
  }
  handleSubmit (event) {
    event.preventDefault();
    axios.post("http://localhost:3333/login", {
      headers: {
        'Access-Control-Allow-Origin': '*',
      },
      proxy: {
        host: "http://localhost",
        port: 3333
      }, 
    }, {
      username: this.state.username,
      password: this.state.password,
    }
      
    )
    .then((res) => console.log(res))
    .catch( (error) => {
      console.log(error)
    })
    
  }
  handleChangeName (e) {
    this.setState({
      username: e.target.value,
      password: this.state.password
    })
  }
  handleChangePassword (e) {
    this.setState({
      username: this.state.username,
      password: e.target.value
    }) 
  }
}