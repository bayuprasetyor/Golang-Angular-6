import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CreateLoaningComponent } from './create-loaning/create-loaning.component';
import { LoaningDetailsComponent } from './loaning-details/loaning-details.component';
import { LoaningListComponent } from './loaning-list/loaning-list.component';
import { HttpClientModule } from '@angular/common/http';
@NgModule({
  declarations: [
    AppComponent,
    CreateLoaningComponent,
    LoaningDetailsComponent,
    LoaningListComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }