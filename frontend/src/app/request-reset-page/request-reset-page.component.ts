import { Component } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { UserService } from '../services/user-service.service';

@Component({
    selector: 'app-request-reset-page',
    templateUrl: './request-reset-page.component.html',
    styleUrls: ['./request-reset-page.component.css']
})
export class RequestResetPageComponent {
    formData = new FormGroup({
        email: new FormControl("")
    })

    constructor(private userService: UserService) { }

    requestReset() {
        if (!this.formData.valid) {
            return
        }

        this.userService.requestResetAuth(this.formData.value.email ?? "").subscribe(data => {
            console.log(data)
            if (data === "202 - Accepted") {
                alert("Check your email!")
            }
        })
    }
}
