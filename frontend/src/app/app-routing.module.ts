import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AccountComponent } from './account/account.component';
import { ChangePassPageComponent } from './change-pass-page/change-pass-page.component';
import { HomePageComponent } from './home-page/home-page.component';
import { LogOutComponent } from './log-out/log-out.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { RequestResetPageComponent } from './request-reset-page/request-reset-page.component';
import { ResetPasswordPageComponent } from './reset-password-page/reset-password-page.component';

const routes: Routes = [
  {path: '', component: HomePageComponent},
  {path: 'register', component: RegisterComponent},
  {path: 'login', component: LoginComponent},
  {path: 'account', component: AccountComponent},
  {path: 'logOut', component: LogOutComponent},
  {path: "requestReset", component: RequestResetPageComponent},
  {path: "resetPass", component: ResetPasswordPageComponent},
  {path: "changePass", component: ChangePassPageComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
