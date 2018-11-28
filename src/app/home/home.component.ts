import { Component, OnInit } from '@angular/core';
import { FileUploader } from 'ng2-file-upload';
import { MatTableDataSource } from '@angular/material';
import { environment } from "../../environments/environment";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  uploader: FileUploader = new FileUploader({
    isHTML5: true,
    autoUpload: true,
    url: `${environment.url}/api/files`,
  });
  hasDropZoneOver: boolean = false;
  displayedColumns: string[] = ['name', 'size', 'progress', 'status']
  dataSource: MatTableDataSource<any> = new MatTableDataSource()

  constructor() {
    this.uploader.onBeforeUploadItem = () => {
      this.dataSource.data = this.uploader.queue
    }
  }
  fileOver(e: any): void {
    this.hasDropZoneOver = e
    this.dataSource.data = this.uploader.queue
  }

  ngOnInit() {
  }

}
