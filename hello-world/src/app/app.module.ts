import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { TestComponent } from './test/test.component';
import { ProfilesComponent } from './profiles/profiles.component';
import { FooterComponent } from './footer/footer.component';
import { HeaderComponent } from './header/header.component';
import { HomeComponent } from './home/home.component';
import { ListingPageComponent } from './listing-page/listing-page.component';
import { SellerPageComponent } from './seller-page/seller-page.component';
import { BComponent} from './login/login.component';
import { SComponent } from './signup/signup.component';

@NgModule({
  declarations: [
    AppComponent,
    TestComponent,
    ProfilesComponent,
    FooterComponent,
    HeaderComponent,
    HomeComponent,
    ListingPageComponent,
    SellerPageComponent,
    BComponent,
    SComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
