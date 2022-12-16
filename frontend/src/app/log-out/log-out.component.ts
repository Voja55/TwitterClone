import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { StoreService } from '../services/store-service.service';

@Component({
  selector: 'app-log-out',
  templateUrl: './log-out.component.html',
  styleUrls: ['./log-out.component.css']
})
export class LogOutComponent implements OnInit {

  constructor(private store: StoreService, private router : Router) { 
    this.logOut()
  }

  ngOnInit(): void {
  }

  logOut() {
    this.store.logout()
    this.router.navigateByUrl("/login");
  }

}
