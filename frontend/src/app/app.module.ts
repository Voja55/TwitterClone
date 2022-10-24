import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

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
    TweetComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgbModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
