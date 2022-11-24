import { Injectable } from '@angular/core';
// import jwt_decode, { JwtPayload } from 'jwt-decode';

@Injectable({
  providedIn: 'root'
})
export class StoreService{

  constructor() { 
    if(sessionStorage.getItem('token')){
      this.loginStatus = true;
      this.token = sessionStorage.getItem('token');
      // this.token = this.getDecodedAccessToken(t);
      // this.role = this.token.role.authority;
      // this.username = this.token.sub;
    }
    else {
      this.loginStatus = false;
    }
  }

  private loginStatus!: boolean;

  private token!: string | null;

  private role : string = "";

  private username : string = "";

  setLoginStatus(status: boolean) {
    this.loginStatus = status;
  }

  setToken(token: string) {
    this.token = token;
    sessionStorage.setItem('jwt', token)
    this.loginStatus = true;
    // this.role = token.role.authority;
    // this.username = token.sub;
    // console.log(this.username);
    // console.log(this.role);
  }

  getToken(): string | null {
    return this.token
  }

  getLoginStatus(): boolean{
    return this.loginStatus
  }

  getDecodedAccessToken(token: any): any {
    try {
      return jwt_decode(token);
    } catch(Error) {
      console.log(Error);
      return null;
    }
  }
}


function jwt_decode(token: any): any {
  console.log('jwt_decode not implemented.');
}

