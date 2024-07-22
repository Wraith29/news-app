import { Component, OnDestroy, OnInit } from '@angular/core';
import { Feed } from '../types/feed';
import { FeedService } from '../services/feed.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrl: './admin.component.css'
})
export class AdminComponent implements OnInit, OnDestroy {
  public feeds: Feed[] = [];

  private _subscriptions: Subscription[] = [];

  constructor(
    private _feedService: FeedService,
  ) { }

  public ngOnInit(): void {
    this._subscriptions.push(this._feedService.getAll().subscribe({
      next: (feeds: Feed[]) => {
        this.feeds = feeds;
      }
    }))
  }

  public ngOnDestroy(): void {
    this._subscriptions.forEach((sub: Subscription) => sub.unsubscribe());
  }
}
