import { LoaningDetailsComponent } from './loaning-details/loaning-details.component';
import { CreateLoaningComponent } from './create-loaning/create-loaning.component';
import { LoaningListComponent } from './loaning-list/loaning-list.component';
import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';


const routes: Routes = [
  { path: '', redirectTo: 'loanings', pathMatch: 'full' },
  { path: 'list', component: LoaningListComponent },
  { path: 'add', component: CreateLoaningComponent },
  { path: 'edit', component: LoaningDetailsComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }