import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';

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
        'Access-Control-Allow-Headers' : 'Content-Type, access-control-allow-methods,access-control-allow-origin,content-type, access-control-allow-headers',
        'Access-Control-Allow-Methods': 'POST, GET, OPTIONS',
        'Access-Control-Allow-Origin' : 'http://localhost:4200/' 
      })
    };
  }

  regUserAuth(username: string, password: string, role: "regular"|"business") {
    return this.client.post<unknown>(environment.apiUrl + "auth_service/users", {
      username: username,
      password: password,
      role: role
    }, this.options())
  }

  loginAuth(username: string, password: string, role: "regular"|"business"){
    return this.client.post<unknown>(environment.apiUrl + "auth_service/login", {
      username: username,
      password: password,
      role: role
    }, this.options())
  }

}
