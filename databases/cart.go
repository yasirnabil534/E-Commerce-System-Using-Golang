package controllers

import ()

var (
	ErrCantFindProduct = errors.New("Product not found")
	ErrCantDecodeProducts = errors.New("Product not found")
	ErrUserIdIsNotValid = errors.New("Invalid user")
	ErrCantUpdateUser = errors.New("")
	ErrCantAddItemCart = errors.New("Cannot add item to cart")
	ErrCantRemoveItemCart = errors.New("Cannot Remove this item from the cart")
	ErrCantGetItem = errors.New("Cart Item not found")
	ErrCantBuyCartItem = errors.New("Cart Item purchase cant update")
)

func AddProductToCart() {}

func RemoveItemFromCart() {}

func BuyItemFromCart() {}

func InstantBuy() {}