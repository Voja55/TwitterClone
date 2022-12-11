import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { environment } from 'src/environments/environment';
import { Profile } from '../model/profile';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  constructor(private client: HttpClient) { }

  options() {
    return  {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        //'Authorization': `Bearer ${sessionStorage.getItem('token')}`,
      })
    };
  }

  getAccount(username: string): Observable<Profile> {
    return this.client.get<Profile>(environment.apiUrl + "profile_service/profile/" + username)
  }

  regUserRegular(username: string, firstName: string, lastName : string, address : string, age : number, gender : boolean) {
    return this.client.post<unknown>(environment.apiUrl + "profile_service/regular", {
      username: username,
      firstName: firstName,
      lastName: lastName,
      address: address,
      age: age,
      gender: gender
    }, this.options())
  }

  regUserBusiness(username: string, email : string, companyName: string, webSite: string) {
    return this.client.post<unknown>(environment.apiUrl + "profile_service/business", {
      username: username,
      email: email,
      companyName: companyName,
      webSite: webSite
    }, this.options())
  }
}
