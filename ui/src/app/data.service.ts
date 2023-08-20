import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { TickData } from 'src/model/tickdata';
import { Config } from '../config/config.service';

@Injectable()
export class DataService {
    constructor(private http: HttpClient) { }

    getData(selectedTimeFrame: string) {
        return this.http.get<TickData>(Config.API_URL + "getTickData?timeFrame=" + selectedTimeFrame);
    }
  }