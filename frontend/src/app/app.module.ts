import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { FooterComponent } from './footer/footer.component';
import { HeaderComponent } from './header/header.component';
import { AComponent} from './home/home.component';
import { BComponent } from './profiles/profiles.component';
import { ListingPageComponent } from './listingPage/listing-page.component';

@NgModule({
  declarations: [
    AppComponent,
    AComponent,
    BComponent,
    FooterComponent,
    HeaderComponent,
    ListingPageComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
