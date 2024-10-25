import React, { useState, useEffect } from "react"
import { getCommentIds, getComment } from "../services/APIService"


/**
 * Comments section component for my website
 */
const CommentsSection = () => {
    const [page, setPage] = useState(1);
    const [commentIds, setCommentIds] = useState([]);
    const [currentComments, setCurrentComments] = useState([]);
    const commentsPerPage = 5;

    useEffect(() => {
        const handleGetCommentIds = async () => {
            try {
                const result = await getCommentIds();
                setCommentIds(result);
            } catch(error) {
                console.error("there was an issue getting the comments: " + error);
            }
        };
        handleGetCommentIds();
    }, []);

    useEffect(() => {
        const getCommentsForPage = async () => {
            const maxCommentsToGet = Math.min(commentIds.length-(page-1)*commentsPerPage,
                commentsPerPage);
            const comments = [];
            for (let i=0; i<maxCommentsToGet; i++) {
                try {
                    const commentId = commentIds[i+(page-1)*commentsPerPage];
                    const comment = await getComment(commentId);
                    comments.push(comment);
                } catch(error) {
                    console.error("Comment could not be found" + error);
                }
            }
            setCurrentComments(comments);
        };
        getCommentsForPage();
    }, [commentIds, page]);

    return (<div> 
        {
            currentComments.map((comment) => {
                return <Comment key={comment.Id} comment={comment}/>;
            })
        }
    </div>);
}

const Comment = ({ comment }) => {
    return (
        <div className="comment-container">
            <p>{comment.Commenter}</p>
            <p>{comment.Content}</p>
            <p>{comment.Timestamp}</p>
        </div>
    );
}

export default CommentsSection