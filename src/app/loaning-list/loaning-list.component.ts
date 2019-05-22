import { LoaningDetailsComponent } from './../loaning-details/loaning-details.component';
import { Observable } from "rxjs";
import { LoaningService } from "./../loaning.service";
import { Loaning } from "./../loaning";
import { Component, OnInit } from "@angular/core";
import { Router } from '@angular/router';

@Component({
  selector: "app-loaning-list",
  templateUrl: "./loaning-list.component.html",
  styleUrls: ["./loaning-list.component.css"]
})

export class LoaningListComponent implements OnInit {
  loanings: Observable<Loaning[]>;
  constructor(private loaningService: LoaningService,private router: Router) {}

  ngOnInit() {
    this.reloadData();
  }

  reloadData() {

   this.loaningService.getLoaningsList().subscribe(data => {
      this.loanings = data;
    })
    console.log(this.loanings);
  }

  deleteLoaning(ID: number) {
    this.loaningService.deleteLoaning(ID)
      .subscribe(
        data => {
          console.log(data);
          this.reloadData();
        },
        error => console.log(error));
  }
  editLoaning(ID: number): void {  
    localStorage.removeItem('editId');  
    localStorage.setItem('editId', ID);  
    this.router.navigate(['edit']);  
  }  
}

