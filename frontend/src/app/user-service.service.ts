import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private client: HttpClient) { }

  regUser(username: string, password: string, role: "regular"|"business") {
    return this.client.post<unknown>(environment.apiUrl + "auth_service/users", {
      username: username,
      password: password,
      role: role
    })
  }
}
