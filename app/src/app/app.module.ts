import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HashLocationStrategy, LocationStrategy } from '@angular/common';

import { AppRoutingModule } from './app-routing.module';
import { HttpClientModule, HttpClient, HttpClientXsrfModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { DatePipe } from '@angular/common'
import {NgxPaginationModule} from 'ngx-pagination';

import { AppComponent } from './app.component';
import { HeaderComponent } from './ui/header/header.component';
import { FooterComponent } from './ui/footer/footer.component';
import { DatabaseComponent } from './page/database/database.component';
import { HeaderDropdownComponent } from './components/dropdown/header-dropdown/header-dropdown.component';
import { RunningComponent } from './page/running/running.component';
import { ScreenComponent } from './page/screen/screen.component';
import { LandingComponent } from './page/landing/landing.component';
 

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    FooterComponent,
    DatabaseComponent,
    HeaderDropdownComponent,
    RunningComponent,
    ScreenComponent,
    LandingComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    NgxPaginationModule,
  ],
  providers: [DatePipe,{provide: LocationStrategy, useClass: HashLocationStrategy}],
  bootstrap: [AppComponent]
})
export class AppModule { }
