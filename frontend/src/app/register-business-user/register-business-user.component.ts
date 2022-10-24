import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-register-business-user',
  templateUrl: './register-business-user.component.html',
  styleUrls: ['./register-business-user.component.css']
})
export class RegisterBusinessUserComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

  user : any = new Object;

  register() {}
  
}
