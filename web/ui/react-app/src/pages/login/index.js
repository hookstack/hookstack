import React, { useState } from 'react';
import * as axios from 'axios';
import ConvoyLogo from '../../assets/img/logo.svg';
import PasswordInvisibleIcon from '../../assets/img/password-visible-icon.svg';
import PasswordVisibleIcon from '../../assets/img/password-invisible-icon.svg';
import { useHistory } from 'react-router-dom';
import { APIURL } from '../../helpers/get-details';
import './style.scss';

const _axios = axios.default;
const request = _axios.create({ baseURL: APIURL });

function LoginPage() {
	const [showLoginPassword, toggleShowLoginPassword] = useState(false);
	const [disableLoginBtn, toggleDisableLoginBtn] = useState(false);
	const [loginDetails, updateLoginDetails] = useState({ username: '', password: '' });
	let route = useHistory();

	const handleUserInput = (event) => {
		const { name, value } = event.target;
		loginDetails[name] = value;
		updateLoginDetails(loginDetails);
	};

	const userLogin = async (event) => {
		event.preventDefault();
		toggleDisableLoginBtn(true);
		try {
			const loginResponse = await (await request({ method: 'POST', url: '/auth/login', data: loginDetails })).data;
			localStorage.setItem('CONVOY_AUTH', JSON.stringify(loginResponse.data));
			toggleDisableLoginBtn(false);
			route.push('/');
		} catch (error) {
			toggleDisableLoginBtn(false);
		}
	};

	return (
		<div className="auth-page">
			<section className="auth-page--container">
				<img src={ConvoyLogo} alt="convoy logo" />

				<form onSubmit={userLogin}>
					<div className="input">
						<label htmlFor="username">Username</label>
						<input type="text" id="username" name="username" autoComplete="username" placeholder="Enter username here" onChange={handleUserInput} />
					</div>

					<div className="input">
						<label htmlFor="password">Password</label>
						<div className="input--password">
							<input type={showLoginPassword ? 'text' : 'password'} id="password" name="password" autoComplete="current-password" placeholder="Enter your password" onChange={handleUserInput} />
							<button className="input--password__view-toggle" type="button" onClick={() => toggleShowLoginPassword(!showLoginPassword)}>
								<img src={showLoginPassword ? PasswordVisibleIcon : PasswordInvisibleIcon} alt={showLoginPassword ? 'hide password icon' : 'view password icon'} />
							</button>
						</div>
					</div>

					<div className="button-container  margin-top">
						<button disabled={disableLoginBtn} className="primary full-100">
							{disableLoginBtn ? 'loading....' : 'Login to Dashboard'}
						</button>
					</div>
				</form>
			</section>
		</div>
	);
}

export { LoginPage };
