import { HttpClient, HttpHeaders } from '@angular/common/http';
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

}
