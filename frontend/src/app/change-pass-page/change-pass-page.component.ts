import { Component } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';
import { Router, ActivatedRoute } from '@angular/router';
import { StoreService } from '../services/store-service.service';
import { UserService } from '../services/user-service.service';

@Component({
  selector: 'app-change-pass-page',
  templateUrl: './change-pass-page.component.html',
  styleUrls: ['./change-pass-page.component.css']
})
export class ChangePassPageComponent {

    constructor(public store : StoreService, private userService: UserService, private router: Router) {

    }

    formData = new FormGroup({
        oldPassword : new FormControl(""),
        newPassword : new FormControl(""),
        repeatPassword : new FormControl("")
      })
    
      changePass() {
        let oldPassword = this.formData.value.oldPassword ?? "";
        let newPassword = this.formData.value.newPassword ?? "";
        let repeatPassword = this.formData.value.repeatPassword ?? "";
        if (!this.formData.valid) {
          return
        }
    
        if (newPassword !== repeatPassword) {
          return
        }

        this.userService.changePassAuth(this.store.getUsername(), oldPassword, newPassword).subscribe(data => {
            this.router.navigateByUrl("/account");
        })
    
      }
}
