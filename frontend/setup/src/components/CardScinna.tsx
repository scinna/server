import React from 'react';
import {Link, Redirect} from 'react-router-dom';

import { withStyles, makeStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import FormControl from '@material-ui/core/FormControl';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Slide from '@material-ui/core/Slide';
import TextField from '@material-ui/core/TextField';
import Tooltip from '@material-ui/core/Tooltip';
import Typography from '@material-ui/core/Typography';

import { useStateValue } from '../context';
import SetScinnaSettings from '../api/Scinna';
import {actionUpdateScinna} from '../actions/scinna';


const useStyles = makeStyles(theme => ({
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
    },
}));

const initialState = {
    ConfigValid: false,
    DialogIpHeaderOpened: false,
};

const Transition = React.forwardRef(function Transition(props, ref) {
    return <Slide direction="up" ref={ref} {...props} />;
});

const HelpTooltip = withStyles(theme => ({
    tooltip: {
      backgroundColor: '#f5f5f9',
      color: 'rgba(0, 0, 0, 0.87)',
      maxWidth: 220,
      fontSize: theme.typography.pxToRem(12),
      border: '1px solid #dadde9',
    },
  }))(Tooltip);
  

export default function() {
    const classes = useStyles();
    const [state, setState] = React.useState(initialState);

    //@ts-ignore
    const [global, dispatch] = useStateValue();

    const handleInputChange = (field: string) => (e: any) => {
        if (e.currentTarget.type === "number") {
            let val = parseInt(e.currentTarget.value)
            dispatch(actionUpdateScinna({ [field]: val }))   
        } else {
            dispatch(actionUpdateScinna({ [field]: e.currentTarget.value }))   
        }
    }

    const handleChange = (field: string) => (event: any) => {
        handleInputChange(field)({ currentTarget: event.target})
    };

    const callAPI = () => {
        SetScinnaSettings(global.Scinna)
            .then((r: any) => {
                setState({
                    ...state,
                    ConfigValid: true,
                    DialogIpHeaderOpened: false,
                })
            })
            .catch((e: any) => {
                console.log(e)
            })
    }

    const submit = (e: any) => {
        e.preventDefault();

        // If there is no IP Header set
        if (global.Scinna.HeaderIPField.length === 0) {
            setState({ ...state, DialogIpHeaderOpened: true })
        } else {
            callAPI();
        }

        return false;
    };

    return <div className="card above">
        { state.ConfigValid ? <Redirect to="/user" /> : null}
        <h4>About this server</h4>
        <form onSubmit={submit}>
            <div className="content centered-form">
                <p>This step is really important. </p>
                <p>Please refer to <a href="https://github.com/scinna/server/wiki/First-launch#scinna-settings" rel="noopener noreferrer" target="_blank">the wiki</a> if you have any issue.</p>
                <FormControl className={classes.formControl} fullWidth>
                    <InputLabel id="registration">Server registration</InputLabel>
                    <HelpTooltip placement="right" title={<React.Fragment>
                            <Typography color="inherit">Private</Typography>
                            <p>No one will be able to register except if you generate <b>Invitation Code</b></p>
                            <Typography color="inherit">Public</Typography>
                            <p>Anyone will be able to register and publish on the server.</p>
                        </React.Fragment>}>
                        <Select labelId="registration" id="registration" value={global.Scinna.RegistrationAllowed} onChange={handleChange("RegistrationAllowed")}>
                            <MenuItem value={"private"}>Private</MenuItem>
                            <MenuItem value={"public"}>Public</MenuItem>
                        </Select>
                    </HelpTooltip>
                </FormControl>
                <HelpTooltip placement="right" title={<React.Fragment>
                            <Typography color="inherit">IP Header</Typography>
                            <p>IP header must be set to the field name that your reverse-proxy sets the client IP.</p>
                            <p>Be sure that the reverse-proxy override this field as an attacker could ban any IP he'd like otherwise.</p>
                            <p>If you are using the wiki's nginx configuration, this should be <code>X-Real-IP</code></p>
                        </React.Fragment>}>
                    <TextField id="scinna_header" label="IP Header" onChange={handleInputChange("HeaderIPField")} value={ global.Scinna.HeaderIPField } fullWidth />
                </HelpTooltip>
                
                <HelpTooltip placement="right" title={<React.Fragment>
                            <Typography color="inherit">Rate limiting</Typography>
                            <p>This settings prevent attacker from doing more than X request per 5 minutes.</p>
                            <p>The default setting is 100 but you might want to change it.</p>
                        </React.Fragment>}>
                    <TextField id="scinna_rate_limit" label="Rate limiting (Per 5 minutes)" required type="number" inputProps={{ min: "0" }} onChange={handleInputChange("RateLimiting")} value={ global.Scinna.RateLimiting } fullWidth />
                </HelpTooltip>
                
                <HelpTooltip placement="right" title={<React.Fragment>
                            <Typography color="inherit">Picture path</Typography>
                            <p>This is where the picture will be stored on the server.</p>
                            <p>Be sure that Scinna has read and write permission on this folder.</p>
                        </React.Fragment>}>
                    <TextField id="scinna_path" required label="Picture path" onChange={handleInputChange("PicturePath")} value={ global.Scinna.PicturePath } fullWidth />
                </HelpTooltip>

                <HelpTooltip placement="right" title={<React.Fragment>
                            <Typography color="inherit">Web URL</Typography>
                            <p>This is the public URL of your Scinna server.</p>
                            <p>It is used for features such as account validation or password recovery.</p>
                        </React.Fragment>}>
                    <TextField id="scinna_url" required label="Web URL" onChange={handleInputChange("WebURL")} value={ global.Scinna.WebURL } fullWidth />
                </HelpTooltip>
            </div>
            <div className="footer">
                <Link className="btn" to="/smtp">Back</Link>
                <input type="submit" className="btn" value="Next" />
            </div>
        </form>
        <Dialog
            open={state.DialogIpHeaderOpened}
            // @ts-ignore
            TransitionComponent={Transition}
            keepMounted
            onClose={() => { setState({ ...state, DialogIpHeaderOpened: false }) } }
            aria-labelledby="alert-dialog-slide-title"
            aria-describedby="alert-dialog-slide-description"
        >
            <DialogTitle id="alert-dialog-slide-title">No IP header field set!</DialogTitle>
            <DialogContent>
                <DialogContentText id="alert-dialog-slide-description">
                    This can be the correct option if you are not using a reverse-proxy, but you're most likely are. <br /><br />
                    Please fill this field since it's extremely important for the security of Scinna and your server. <br /><br />
                    Informations can be found on the <a href="https://github.com/scinna/server/wiki/First-launch#scinna-settings" rel="noopener noreferrer" target="_blank">wiki</a>.
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={() => { setState({ ...state, DialogIpHeaderOpened: false }) }} color="primary">
                    Cancel
                </Button>
                <Button onClick={() => { callAPI() } } color="primary">
                    Next
                </Button>
            </DialogActions>
      </Dialog>
    </div>;
}