/* eslint-disable react/prefer-stateless-function */
import React, { Component }  from 'react';
import MaskedInput from 'react-text-mask';
import Input from 'material-ui/Input';
import PropTypes from 'prop-types';
import Card from 'material-ui/Card';
import { withStyles } from 'material-ui/styles';
import Button from 'material-ui/Button';
import SearchIcon from 'material-ui-icons/Search';

const styles = theme => ({
    container: {
        display: 'inline',
    },
    card: {
        paddingLeft: 0,
        paddingBottom: 0,
        marginBottom: 10,
        width: "100%",
        display: 'flex',
        float: "left"
    },
    errorCard: {
        backgroundColor: "#f2dede",
        borderColor: "#ebcccc",
        color: "#a94442",
        width: "100%",
        display: 'flex',
        marginTop: 20,
        paddingTop: 10,
        paddingBottom: 10,
    },
    errorText: {
        paddingLeft: 20
    },
    inputBox: {
        paddingLeft: 20,
        fontSize: "1rem",
        display: 'inline',
        width: "100%",
    },
    button: {
        float: "right",
        backgroundColor: "#436280"
    },
    textField : {
        marginLeft: 20,   
        fontSize: "1rem",
    }
});

class PhoneSearchBar extends Component {
    constructor(props) {
        super(props);
        this.state = {
            value: '',
            name: "",
            error: false,
            multipleHits: false
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }
  
    handleChange = name => event => {
        this.setState({
            [name]: event.target.value,
        });
    };

    handleSubmit(event) {
        //Send the masked phonenumber lookup request to the parent
        var maskedNum = this.state.value.replace(/\D/g,'');

        // Confirm correct length
        if(maskedNum.length !== 10){
            this.setState({error : true});
        }
        else
        {
            this.props.lookupNumber(maskedNum, this.state.name);
            this.setState({error : false});
        }
        
        // Prevent default button behavior (reloading of DOM)
        event.preventDefault();
    }
  
    render() {
      const { classes } = this.props;
      return (
        <div className={classes.container}>
            <Card className={classes.card}>
                <MaskedInput 
                    className={classes.inputBox}
                    mask={['(', /[1-9]/, /\d/, /\d/, ')', ' ', /\d/, /\d/, /\d/, '-', /\d/, /\d/, /\d/, /\d/]}
                    placeholder="e.g 908-826-6380"
                    guide={true}
                    value={this.state.value} 
                    id="my-input-id"
                    onChange={this.handleChange('value')} 
                />
                <Button raised className={classes.button} onClick={this.handleSubmit}>
                    <SearchIcon color="contrast"/>
                </Button>
            </Card>
            {this.state.error &&
                <Card className={classes.errorCard}>
                    <div className={classes.errorText}>Full phone with area code required (555-555-1234), U.S. numbers only.</div>
                </Card>
            }
            {this.props.response === -2 &&
                <Card className={classes.errorCard}>
                    <div className={classes.errorText}>We could not find any records for that search criteria.</div>
                </Card>
            }
            {(this.props.response === -3 || this.state.name !== "") &&
                <div>
                    <Card className={classes.card}>
                        <Input
                            id="name"
                            fullWidth
                            placeholder="Multiple hits found. What's the name associated with the number?"
                            className={classes.textField}
                            disableUnderline
                            value={this.state.name}
                            onChange={this.handleChange('name')}
                        />
                    </Card>
                </div>
            }
            {this.props.response === -4 &&
                <Card className={classes.errorCard}>
                    <div className={classes.errorText}> Serverside error. Could not lookup information</div>
                </Card>
            }
        </div>
      );
    }
  }
  
  PhoneSearchBar.propTypes = {
    classes: PropTypes.object.isRequired,
  };
  
  export default withStyles(styles)(PhoneSearchBar);
