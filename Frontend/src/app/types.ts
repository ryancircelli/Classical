export type Class = {
  upvotes: number;
  downvotes: number;
  className: string;
  total_votes: number;
  lastUpdated: string;
}

export type Post = {
  postId: number;
  className: string;
  postVotes: number;
  postName: string;
  postContent: string;
  upvoted: boolean;
  downvoted: boolean;
  timePosted: string;
}