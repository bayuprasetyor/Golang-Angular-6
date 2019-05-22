import { Loaning } from './../loaning';
import { Component, OnInit, Input } from '@angular/core';
import { LoaningService } from '../loaning.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-loaning-details',
  templateUrl: './loaning-details.component.html',
  styleUrls: ['./loaning-details.component.css']
})
export class LoaningDetailsComponent implements OnInit {

	loaning: Loaning = new Loaning();

	submitted = false;
	constructor(private loaningService: LoaningService,private router: Router) { }

	ngOnInit() {
		this.reloadData();
	}


	reloadData() {
		let id = localStorage.getItem('editId');  
		if (+id > 0) {  
		  this.loaningService.getLoaning(+id).subscribe(data => {  
		    this.loaning = data;
		    this.submitted = false;
		  })  
		}  
		console.log(this.loaning);
  }
	save(ID: number) {
		this.loaningService.updateLoaning(ID,this.loaning)
		  .subscribe(data => console.log(data), error => console.log(error));
		this.loaning = new Loaning();
	}

	onUpdate(ID: number) {
		this.submitted = true;
		this.save(ID);
		this.router.navigate(['list']);  
	}
}
