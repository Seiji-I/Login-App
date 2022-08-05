import React from "react";
import axios from "axios";
import { Link } from "react-router-dom";

export default class Top extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      username: ""
    }
    this.getLoginUser = this.getLoginUser.bind(this);
    this.getLoginUser()
  }
  render() {
    let element;
    if (this.state.username != null) {
      element = <>
        <p>{`Hello ${this.state.username}`}</p>
        <p><Link to="/logout">Logout</Link></p>
      </>
    } else {
      element = <p><Link to="/login">Login</Link></p>
    }
    return (
      <>
        <h1>Welcome Top page!</h1>
        <div>{element}</div>
      </>
    )
  }
  getLoginUser() {
    return axios.get("http://127.0.0.1:3333/login", {
      proxy: {
        host: "http://localhost",
        port: 3333
      },
    })
    .then(response => {
      const username = response.data.user;
      this.setState({ username: username});
      setTimeout(() => console.log(username), 3000);
    })
    .catch(error => {
      console.log(error)
    })
  }
}