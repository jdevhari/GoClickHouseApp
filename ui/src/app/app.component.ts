import { Component, OnInit } from '@angular/core';
import type { EChartsOption } from 'echarts';
import { DataService } from './data.service';
import { TickData } from 'src/model/tickdata';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [DataService]
})
export class AppComponent implements OnInit {
  title = 'ui';
  instrumentSymbol = 'NIFTY';
  dataAvailable = false;
  selectedTimeFrame = "1";
  options: EChartsOption = {};

  constructor(private service: DataService) { }
  
  handleChange(tf: any){
    this.selectedTimeFrame = tf;
    this.process(this.selectedTimeFrame);
  }

  ngOnInit(): void {
    this.process(this.selectedTimeFrame);
  }

  process(selectedTimeFrame : string): void {
    this.service.getData(selectedTimeFrame)
      .subscribe((displayData: TickData) => {
        this.options = {
          xAxis: {
            data: displayData.dateTime
          },
          yAxis: {min:17500, max:20500},
          series: [
            {
              type: 'candlestick',
              data: displayData.tickData
            }
          ]
        }
        this.dataAvailable = true;
      },err => {
        console.log("Error caught at Subscriber " + err)
      }
      );
  }  
}
