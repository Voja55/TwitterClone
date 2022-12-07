import { Injectable } from '@angular/core';
import jwt_decode from 'jwt-decode';

@Injectable({
  providedIn: 'root'
})
export class StoreService{

  constructor() { 
    if(sessionStorage.getItem('jwt')){
      var t = sessionStorage.getItem('jwt')
      this.loginStatus = true;
      this.decodedToken = this.getDecodedAccessToken(t)
      this.role = this.decodedToken.role;   
      this.username = this.decodedToken.username;
      console.log(this.decodedToken)
      console.log(this.username)
      console.log(this.role)
    }
    else {
      this.loginStatus = false;
    }
  }

  private loginStatus!: boolean;

  private token!: string | null;

  private decodedToken : any;

  private role : string = "";

  private username : string = "";

  getToken(): string | null {
    return this.token
  }

  getLoginStatus(): boolean{
    return this.loginStatus
  }

  getUsername(): string{
    return this.username
  }

  getRole(): string {
    return this.role
  }

  login(token: string) {
    sessionStorage.setItem('jwt', token)
    this.token = token;
    
    this.loginStatus = true;
    this.decodedToken = this.getDecodedAccessToken(this.token);
    this.role = this.decodedToken.role;
    this.username = this.decodedToken.username;

    console.log(this.username);
    console.log(this.role);
  }

  logout() {
    sessionStorage.removeItem('jwt')
    this.loginStatus = false;
    this.decodedToken = null;
    this.role = ""
    this.username = ""
    this.token = null   
  }

  getDecodedAccessToken(token: any): any {
    try {
      console.log(jwt_decode(token))
      return jwt_decode(token);
    } catch(Error) {
      console.log(Error);
      return null;
    }
  }
}