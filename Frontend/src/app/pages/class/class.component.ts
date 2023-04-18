import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router'
import { ClassAPIService } from 'src/app/services/class.services';
import { Post } from 'src/app/types';
import { isWebUri } from 'valid-url';

@Component({
  selector: 'app-class',
  templateUrl: './class.component.html',
  styleUrls: ['./class.component.css']
})
export class ClassComponent {

  constructor(private route: ActivatedRoute, public classAPIService: ClassAPIService) { }

  class: string = ""

  posts: Post[] = []

  newPost: string = "";
  errorMessage: string = "";

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.class = params.get('className') ?? "";
      this.loadData();
    });
  }

  rankPosts(results: Post[]) {
    results.sort((a, b) => b.postVotes - a.postVotes);
    return results;
  }

  submitPost() {
    return new Promise((resolve, reject) => {
      if (this.newPost === "") {
        this.errorMessage = 'Please provide a link!';
        this.newPost = "";
        reject(this.errorMessage);
        return;
      } else if (isWebUri(this.newPost) === undefined) {
        this.errorMessage = 'Invalid Link!\nMake sure to include http:// or https://';
        this.newPost = "";
        reject(this.errorMessage);
        return;
      }

      this.classAPIService.createClassPost(this.class, this.newPost)
        .subscribe(
            response => {
              this.loadData();
            },
            error => {
            }
          );
      this.newPost = "";
      this.errorMessage = ""
      resolve("");
    });

  }

  upvote(id: number) {
    if (this.posts[id].downvoted) {
      this.classAPIService.increasePostVotes(this.class, this.posts[id].postId).subscribe(
        response => {},
        error => {}
      );
      this.posts[id].postVotes = this.posts[id].postVotes + 1;
    }
    this.classAPIService.increasePostVotes(this.class, this.posts[id].postId).subscribe(
      response => {},
      error => {}
    );
    this.posts[id].upvoted = true;
    this.posts[id].downvoted = false;
    this.posts[id].postVotes = this.posts[id].postVotes + 1;
  }

  resetVote(id : number){
    if (this.posts[id].upvoted) {
      this.classAPIService.decreasePostVotes(this.class, this.posts[id].postId).subscribe(
        response => {},
        error => {}
      );
      this.posts[id].postVotes = this.posts[id].postVotes - 1;
    }
    if (this.posts[id].downvoted) {
      this.classAPIService.increasePostVotes(this.class, this.posts[id].postId).subscribe(
        response => {},
        error => {}
      );
      this.posts[id].postVotes = this.posts[id].postVotes + 1;
    }
    this.posts[id].upvoted = false;
    this.posts[id].downvoted = false;
  }

  downvote(id : number) {
    if (this.posts[id].upvoted) {
      this.classAPIService.decreasePostVotes(this.class, this.posts[id].postId).subscribe(
        response => {},
        error => {}
      );
      this.posts[id].postVotes = this.posts[id].postVotes - 1;
    }
    this.classAPIService.decreasePostVotes(this.class, this.posts[id].postId).subscribe(
      response => {},
      error => {}
    );
    this.posts[id].postVotes = this.posts[id].postVotes - 1;
    this.posts[id].upvoted = false;
    this.posts[id].downvoted = true;
  }

  loadData() {
    console.log("reloading data")
    this.posts = []
    this.classAPIService.getClassPosts(this.class).subscribe(data => {
      this.posts = this.rankPosts(data.map(postData => ({
        ...postData,
        timePosted: new Date(parseInt(postData.timePosted) * 1000).toLocaleString()
      })));
    }); 
  }
  
}
