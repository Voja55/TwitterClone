import { Component, OnInit } from '@angular/core';
import { User } from '../model/user';
import { StoreService } from '../services/store-service.service';
import { UserService } from '../services/user-service.service';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.css']
})
export class AccountComponent implements OnInit {

  constructor(private store : StoreService, private userService : UserService) { }

  ngOnInit(): void {
  }

  user : User = new User();
  statusDisplayName : boolean = false;
  statusDescription : any;
  statusPassword : any;
  changeDescription(){}
  changeDisplayName(){}
  changePassword(){}

  confirm() {
    console.log(this.store.getUsername());
    this.user.username = this.store.getUsername();
    let codeField = document.getElementById("code") as HTMLInputElement;
    this.userService.confirmAuth(this.user.username, +codeField.value).subscribe(data => {
      console.log(data);
    })
  }
}
