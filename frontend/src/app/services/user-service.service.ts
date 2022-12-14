import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Token } from '@angular/compiler';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { Jwt } from '../model/jwt';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private client: HttpClient) { }

  options() {
    return  {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        //'Authorization': `Bearer ${sessionStorage.getItem('token')}`,
      })
    };
  }

  regUserAuth(username: string, password: string, email : string, role: "regular"|"business") {
    return this.client.post<unknown>(environment.apiUrl + "auth_service/users", {
      username: username,
      password: password,
      email : email,
      role: role
    }, this.options())
  }

  loginAuth(username: string, password: string, role: "regular"|"business"){
    return this.client.post<Jwt>(environment.apiUrl + "auth_service/login", {
      username: username,
      password: password,
      role: role
    }, this.options())
  }

  confirmAuth(username: string, code: number){
    return this.client.post<Jwt>(environment.apiUrl + "auth_service/confirm", {
      username: username,
      ccode: code
    }, this.options())
  }

  requestResetAuth(email: string){
    return this.client.post<unknown>(environment.apiUrl + "auth_service/requestreset", {
      email: email
    }, this.options())
  }

  resetPassAuth(token: string, password: string){
    return this.client.post<unknown>(environment.apiUrl + "auth_service/reset", {
      token: token,
      password: password
    })
  }

  changePassAuth(username: string, oldPassword: string, newPassword: string){
    return this.client.post<unknown>(environment.apiUrl + "auth_service/changepass", {
        username: username,
        oldPassword: oldPassword,
        newPassword: newPassword
    })
  }

  resendCCodeAuth(username: string){
    return this.client.post<unknown>(environment.apiUrl + "auth_service/resend", {
        username: username
    })
  }

}
