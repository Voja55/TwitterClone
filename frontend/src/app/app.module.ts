import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { HomePageComponent } from './home-page/home-page.component';
import { RegisterComponent } from './register/register.component';
import { LoginComponent } from './login/login.component';
import { AccountComponent } from './account/account.component';
import { RegisterUserComponent } from './register-user/register-user.component';
import { RegisterBusinessUserComponent } from './register-business-user/register-business-user.component';
import { TweetsComponent } from './tweets/tweets.component';
import { TweetPopupComponent } from './tweet-popup/tweet-popup.component';
import { TweetComponent } from './tweet/tweet.component';
import { HttpClientModule } from '@angular/common/http';
import { UserService } from './services/user-service.service';
import { LogOutComponent } from './log-out/log-out.component';
import { RecaptchaModule } from 'ng-recaptcha';
import { RequestResetPageComponent } from './request-reset-page/request-reset-page.component';
import { ResetPasswordPageComponent } from './reset-password-page/reset-password-page.component';

@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    RegisterComponent,
    LoginComponent,
    AccountComponent,
    RegisterUserComponent,
    RegisterBusinessUserComponent,
    TweetsComponent,
    TweetPopupComponent,
    TweetComponent,
    LogOutComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgbModule,
    FormsModule,
    HttpClientModule,
    RecaptchaModule,
    ReactiveFormsModule
  ],
  providers: [UserService],
  bootstrap: [AppComponent]
})
export class AppModule { }
