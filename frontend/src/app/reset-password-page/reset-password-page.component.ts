import { Component } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { UserService } from '../services/user-service.service';

@Component({
  selector: 'app-reset-password-page',
  templateUrl: './reset-password-page.component.html',
  styleUrls: ['./reset-password-page.component.css']
})
export class ResetPasswordPageComponent {

  constructor(private userService: UserService, private router: Router, private route: ActivatedRoute) {

  }

  formData = new FormGroup({
    password : new FormControl(""),
    repeatPassword : new FormControl("")
  })

  changePass() {
    let password = this.formData.value.password ?? "";
    let repeatPassword = this.formData.value.repeatPassword ?? "";
    if (!this.formData.valid) {
      return
    }

    if (password !== repeatPassword) {
      return
    }

    this.route.queryParams.subscribe(params => {
      let token = params['id'];
      this.userService.resetPassAuth(token, password).subscribe( data => {
        this.router.navigateByUrl("/login");
      })
    })

    
  }
}
