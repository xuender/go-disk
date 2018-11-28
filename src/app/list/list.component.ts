import { Component, OnInit, ViewChild } from '@angular/core';
import { MatSort, MatTableDataSource, MatDialog } from '@angular/material';
import { Observable } from 'rxjs';
import { indexOf } from 'lodash'

import { File, Type } from '../file';
import { FilesService } from '../files.service';
import { SingleComponent } from '../photo/single/single.component';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit {
  dataSource: MatTableDataSource<File> = new MatTableDataSource()
  displayedColumns: string[] = ['name', 'type', 'ca']
  paths = []
  @ViewChild(MatSort) sort: MatSort;
  constructor(
    private filesServer: FilesService,
    public dialog: MatDialog,
  ) {
  }

  dir(dir: string) {
    if (dir) {
      this.paths.push(dir)
    }
    this.list()
  }
  back(dir = '') {
    if (dir) {
      const i = indexOf(this.paths, dir)
      if (i > -1) {
        this.paths = this.paths.slice(0, i + 1)
      }
    } else {
      this.paths = []
    }
    this.list()
  }
  list() {
    this.filesServer.files$(this.paths.join('/')).subscribe(data => this.dataSource.data = data)
  }

  ngOnInit() {
    this.dataSource.sort = this.sort;
    this.list()
  }

  open(file: File) {
    if (file.type == Type.DIR) {
      this.dir(file.name)
    } else {
      this.dialog.open(SingleComponent, {
        data: file,
      });
    }
  }
}
