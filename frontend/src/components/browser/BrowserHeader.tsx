import {IconButton, TextField} from "@material-ui/core";
import {
    ArrowBack as BackButton,
    ArrowForward as ForwardButton,
    Refresh
} from '@material-ui/icons';

export const BrowserHeader = () => {
    return <div>
        <IconButton>
            <BackButton/>
        </IconButton>
        <IconButton>
            <ForwardButton/>
        </IconButton>
        <IconButton>
            <Refresh/>
        </IconButton>
        <TextField/>
    </div>;
}