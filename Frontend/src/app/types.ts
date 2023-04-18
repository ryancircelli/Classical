export type Class = {
  upvotes: number;
  downvotes: number;
  className: string;
  dateUpdated: Date;
  total_votes: number;
}

export type Post = {
  postID: number;
  className: string;
  postVotes: number;
  postName: string;
  postContent: string;
  upvoted: boolean;
  downvoted: boolean;
}