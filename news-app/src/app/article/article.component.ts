import { Component, Input } from "@angular/core";
import { Article } from "../types/article";

@Component({
  selector: "app-article",
  templateUrl: "./article.component.html",
  styleUrl: "./article.component.css",
})
export class ArticleComponent {
  @Input({ required: true })
  public article!: Article;
  public collapsed: boolean = true;

  public toggleCollapsed(): void {
    this.collapsed = !this.collapsed;
  }
}
