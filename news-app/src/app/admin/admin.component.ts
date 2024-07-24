import { Component, OnDestroy, OnInit } from '@angular/core';
import { Feed } from '../types/feed';
import { FeedService } from '../services/feed.service';
import { Subscription } from 'rxjs';
import { ConfirmationService } from 'primeng/api';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrl: './admin.component.css',
})
export class AdminComponent implements OnInit, OnDestroy {
  public feeds: Feed[] = [];

  private _subscriptions: Subscription[] = [];
  private _tempFeedStore: { [id: number]: Feed } = {};

  constructor(private _feedService: FeedService, private _confirmationService: ConfirmationService) { }

  public ngOnInit(): void {
    this._subscriptions.push(
      this._feedService.getAll().subscribe({
        next: (feeds: Feed[]) => {
          this.feeds = feeds;
        },
      }),
    );
  }

  public ngOnDestroy(): void {
    this._subscriptions.forEach((sub: Subscription) => sub.unsubscribe());
  }

  public deleteFeed(feed: Feed): void {
    this._confirmationService.confirm({
      header: 'Delete Feed',
      message: 'Are you sure you want to delete this feed?',
      accept: () => {
        this._feedService.delete(feed.id);
      },
      acceptButtonStyleClass: 'p-button-danger'
    });

  }

  public onFeedEditInit(feed: Feed): void {
    // Create a temporary copy of the item before editing - to restore if editing is cancelled
    this._tempFeedStore[feed.id] = { ...feed };
  }

  public onFeedEditSave(feed: Feed): void {
    this._feedService.update(feed);
  }

  public onFeedEditCancel(feed: Feed, index: number): void {
    this.feeds[index] = this._tempFeedStore[feed.id];
    delete this._tempFeedStore[feed.id];
  }
}
