import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.css']
})
export class AccountComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

  user : any = new Object;
  statusDisplayName : boolean = false;
  statusDescription : any;
  statusPassword : any;
  changeDescription(){}
  changeDisplayName(){}
  changePassword(){}

}
