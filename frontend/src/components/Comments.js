import React, { useState, useEffect, useCallback } from "react"
import { getCommentIds, getComment, postComment } from "../services/APIService"
import "../styles/Comments.css"

/**
 * Comments section component for my website
 */
const CommentsSection = () => {
    const [page, setPage] = useState(1); // controls pagination
    const [commentIds, setCommentIds] = useState([]); // the list of comments returned by the database
    const [currentComments, setCurrentComments] = useState([]); // comments displayed on the current page
    const [showHiddenBen, setShowHiddenBen] = useState(false); // whether or not a silly me is showing!
    const [commentPostResponse, setCommentPostResponse] = useState(""); // the status of the last post made
    const commentsPerPage = 5;

    // used to get the list of comments from the server. called on page load and posted comment
    const handleGetCommentIds = async () => {
            try {
                const result = await getCommentIds();
                setCommentIds(result);
            } catch(error) {
                console.error("there was an issue getting the comments: " + error);
            }
        }

    useEffect(() => {
        handleGetCommentIds();
    }, []);

    // whenever the list of comment ids or the page number is updated, we should re-fetch the relevant comments
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
     * method to handle the posting of comments.
     * @param {*} formData the input fields on the form 
     */
    function handleComment(formData) {
        formData.preventDefault();
        const username = formData.target.username.value;
        const content = formData.target.content.value;
        if (username == "" || content == "" ) {
            setCommentPostResponse("Please fill out all fields!");
            return;
        }
        const timestamp = new Date().toISOString();

        postComment(username, content, timestamp).then(response => {
           setCommentPostResponse(response); 
        }); 
        handleGetCommentIds();

        formData.target.content.value = "";
    }

    // The actual content of this module is in 3 sections: posting, comments, navigation
    return (<div className="comment-section"> 

        {/*this section handles the posting of comments*/}
        <div className="comment-creator">
            <form onSubmit={handleComment}>
                <div className="comment-creator-row">
                    <label for="username">username: </label>
                    <input type="text" id="username" name="username"/>
                </div>
                <div className="comment-creator-row">
                    <textarea type="text" id="comment-content" name="content" style={{flex: 8}} placeholder="your comment here!"/>
                    <input type="submit" value="post" style={{flex:1}} id="comment-post" class="section"/>
                </div>
            </form>
            <div>{commentPostResponse}</div>
        </div>
        {/*this section renders the comments that are on the current page*/}
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
        {/* this section handles the pagination and navigation */}
        <div id="comment-navigation">
            <a onClick={() => {setPage(Math.max(page-1, 1))}} style={{cursor: "pointer"}} id="comment-navigation-prev">prev</a>
            {
                [...Array(Math.ceil(commentIds.length/commentsPerPage)).keys()].map((index) => {
                    return <div onClick={() => setPage(index+1)} style={{cursor: "pointer"}}>{index+1}</div>
                })
            }
            <a onClick={() => {setPage(Math.min(page+1, Math.ceil(commentIds.length/commentsPerPage)));}}
                style={{cursor: "pointer"}} id="comment-navigation-next">next</a>
        </div>
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