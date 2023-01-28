import { Component} from '@angular/core';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {

  switchForm() {
    this.register = !this.register;
  }

  register : boolean = true;

}
