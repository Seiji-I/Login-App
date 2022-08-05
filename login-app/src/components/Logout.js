import React from "react";
import axios from "axios";
import { Link } from "react-router-dom";

export default class Logout extends React.Component {
  constructor(props) {
    super(props)
    this.state = {}
    this.handleLogout = this.handleLogout.bind(this)
    this.handleLogout()
  }
  render() {
    return (
      <>
        <p>This page is user logout</p>
        <p><Link to="/">Top Page</Link></p>
      </>
    )
    
  }
  handleLogout () {
    axios.get("http://localhost:3333/logout",{
      headers: {
        'Access-Control-Allow-Origin': '*',
      },
      proxy: {
        host: "http://localhost",
        port: 3333
      },
      withCredentials: true,
    }
      
    )
  }
}