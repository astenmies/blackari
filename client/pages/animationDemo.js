
import React, { Component } from 'react'
import styled from 'styled-components'
import { TimelineMax } from 'gsap'
import Cat from '../components/github-logo.svg'

const MyCat = styled(Cat)`
  width: 50px;
  height: 50px;
  margin-top: 40px;
`;

const Box = styled.div`
  width: 150px;
  padding: 20px;
  background: white;
  border-radius: 4px;
  box-shadow: rgba(0, 0, 0, .3) 0 2px 4px;
  text-align: center;
  overflow: hidden;
`;

const List = styled.ul`
  position: relative;
  list-style-type: none;
  margin: 40px 0 0 0;
  padding: 0;
`;

const ListItem = styled.li`
  position: relative;
  left: 0;
  right: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 0;
  width: 130px;
  background: #bb024f;
  border-radius: 3px;
  color: white;
  opacity: 0;
  transform: translateX(-200px);
`;

const Button = styled.button`
  display: block;
  width: 100%;
  margin: 10px 0 0;
  padding: 10px;
  border: 0;
  border-radius: 3px;
  background: linear-gradient(45deg, #f63954, #0c5e8e);
  color: white;
  text-align: center;
  cursor: pointer;
`;
const tl = new TimelineMax();
export default class Hire extends React.Component{
  constructor(props){
    super(props);
    this.state = {
        done: false
    }
  }
  handleClick = () => {
    let anim = ({onComplete, isReverse}) => {
        return (
            tl.to('button', 0.5, { y: 30, opacity: 0, visibility: 'hidden', ease: Power4.easeOut, onComplete: () => { if(isReverse){onComplete} } })
            .to('#box', 1, {scale: 1.1, ease: Elastic.easeOut.config(1, 0.3)})
            .to('li', 1, { height: '20px', margin: '5px auto', padding: '10px', ease: Elastic.easeOut.config(1, 0.3) }, '-=1')
            .to('#svg', 1, {y: '-40px', ease: Bounce.easeOut}, '-=1')
            .staggerTo('li', 0.2, { x: 0, y: '-10px', fontSize: '10px', opacity: 1 }, 0.1)
            .to('ul', 1, {scaleY: 0.9, y: '-10px', onComplete: () => { if(!isReverse){onComplete} } }, "-=0.5")
        )
    }
    if(!this.state.done) {
        anim({onComplete: this.setState({done: !this.state.done}), isReverse: false}).play()
    } else {
        anim({onComplete: this.setState({done: !this.state.done}), isReverse: true}).reverse(2, false)
    }
  }
  
  render(){
      console.log(this.state)
    return(
      <div style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '100vh',
        backgroundColor: '#f63954'
        }}>
          <Box id="box">
            <div onClick={this.handleClick} style={{cursor: 'pointer'}} >
                <MyCat id="svg" />
            </div>
            <List>
            <ListItem>One</ListItem>
            <ListItem>Two</ListItem>
            <ListItem>Three</ListItem>
            <ListItem>Four</ListItem>
            <ListItem>Five</ListItem>
            </List>
            <Button onClick={this.handleClick}>Start animation</Button>
        </Box>
      </div>
      );
  }
}