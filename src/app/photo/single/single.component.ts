import { Component, OnInit, Inject } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material';
import { DatePipe } from '@angular/common';
import { File } from 'src/app/file';
import * as crypto from 'crypto-js'
import { environment } from "../../../environments/environment";

@Component({
  selector: 'app-single',
  templateUrl: './single.component.html',
  styleUrls: ['./single.component.scss']
})
export class SingleComponent implements OnInit {
  src: string
  photo: File
  constructor(
    @Inject(MAT_DIALOG_DATA) private data: File,
    private dataPipe: DatePipe,
  ) {
    this.photo = data
    this.src = `${environment.url}/photo/${
      dataPipe.transform(data.ca, 'yyyy/MM/dd')
      }/${crypto.enc.Hex.stringify(crypto.enc.Base64.parse(data.id))}.jpg`
  }

  ngOnInit() {
  }

}
