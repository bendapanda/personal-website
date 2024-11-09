import React, { useState, useEffect } from "react"
import { getCommentIds, getComment, postComment } from "../services/APIService"
import "../styles/Comments.css"

/**
 * Comments section component for my website
 */
const CommentsSection = () => {
    const [page, setPage] = useState(1);
    const [commentIds, setCommentIds] = useState([]);
    const [currentComments, setCurrentComments] = useState([]);
    const [showHiddenBen, setShowHiddenBen] = useState(false);
    const [commentPostResponse, setCommentPostResponse] = useState("");
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

    /**
     * method to handle the posting of comments 
     * @param {*} formData the input fields on the form 
     */
    function handleComment(formData) {
        formData.preventDefault();
        const username = formData.target.username.value;
        const content = formData.target.content.value;
        const timestamp = Date.UTC();

        postComment(username, content, timestamp).then(response => {
           setCommentPostResponse(response); 
        }); 

    }

    return (<div className="comment-section"> 
        <div className="comment-creator">
            <form onSubmit={handleComment}>
                <label for="username">username: </label>
                <input type="text" id="username" name="username"/>
                <textarea type="text" id="content" name="content"/>
                <input type="submit" value="post"/>
            </form>
            <div>{commentPostResponse}</div>
        </div>
        {
            currentComments.map((comment, index) => {
                // As a fun feature, We add a secret me that appears on hover on the second comment
                if (index == 1) {
                    return (
                        <div style={{position: "relative"}}>
                            {
                                showHiddenBen &&
                                <img id="hidden-ben-comments" src={`${process.env.PUBLIC_URL}/resources/ben_squat.png`}/>
                            }
                            <div style={{zIndex: "1"}} onMouseEnter={() => setShowHiddenBen(true)} onMouseLeave={() => setShowHiddenBen(false)}>
                                <Comment key={comment.Id} comment={comment} layer={2}/>
                            </div>
                        </div>
                    );
                }
                return <Comment key={comment.Id} comment={comment} layer={0}/>;
            })
        }
    </div>);
}



/**
 *  returns a rendering of one comment object. 
 * @param comment the comment object to be rendered
 * @returns 
 */
const Comment = ({ comment, layer }) => {
    return (
        <div className="comment-container" style={{position: "relative", zIndex: layer}}>
            <div className="comment-info">
                <h3 className="comment-commenter">{comment.Commenter}</h3>
                <h4 className="comment-date">{new Date(comment.Timestamp).toDateString()}</h4>
            </div>
            <p className="comment-content">{comment.Content}</p>
        </div>
    );
}

export default CommentsSection