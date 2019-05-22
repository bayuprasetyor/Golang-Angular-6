import { LoaningService } from './../loaning.service';
import { Loaning } from './../loaning';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-create-loaning',
  templateUrl: './create-loaning.component.html',
  styleUrls: ['./create-loaning.component.css']
})
export class CreateLoaningComponent implements OnInit {

  loaning: Loaning = new Loaning();
  submitted = false;

  constructor(private loaningService: LoaningService,private router: Router) { }

  ngOnInit() {
  }

  newLoaning(): void {
    this.submitted = false;
    this.loaning = new Loaning();
  }

  save() {
    this.loaningService.createLoaning(this.loaning)
      .subscribe(data => console.log(data), error => console.log(error));
    this.loaning = new Loaning();
  }

  onSubmit() {
    this.submitted = true;
    this.save();
    this.router.navigate(['list']);
  }
}