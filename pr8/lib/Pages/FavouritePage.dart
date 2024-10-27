import 'package:flutter/material.dart';
import 'package:pr8/Pages/ItemPage.dart';
import 'package:pr8/api_service.dart';
import 'package:pr8/model/items.dart';

class FavoritePage extends StatefulWidget {
  const FavoritePage(
      {super.key, required this.updateCount, required this.navToShopCart});

  final Function() updateCount;
  final Function(int i) navToShopCart;

  @override
  State<FavoritePage> createState() => _FavoritePageState();
}

class _FavoritePageState extends State<FavoritePage> {
  late Future<List<Items>> ItemsFavList;

  @override
  void initState() {
    super.initState();
    ItemsFavList = ApiService().getFavoriteProducts();
  }

  /*
  int findIndexById(int id) {
    return ItemsList.indexWhere((item) => item.id == id);
  }

  void AddFavorite(int index) {
    setState(() {
      if (!Favorite.any((el) => el == index)) {
        Favorite.add(index);
      } else {
        Favorite.remove(index);
      }
      ItemsFavList = ItemsList.where(
          (item) => Favorite.any((element) => element == item.id)).toList();
    });
  }*/

  void NavToItem(index) async {
    int? answ = await Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => ItemPage(
          index: index,
          updateCount: () => widget.updateCount(),
          navToShopCart: (i) => widget.navToShopCart(i),
        ),
      ),
    );
    /*
    setState(() {
      if (answ != null) {
        ItemsList.removeAt(answ);
      }
      ItemsFavList = ItemsList.where(
          (item) => Favorite.any((element) => element == item.id)).toList();
    });
    */
  }

/*
  void AddShopCart(index) async {
    setState(() {
      if (!ShoppingCart.any((el) => el.id == index)) {
        ShoppingCart.add(ShoppingCartItem(index, 1));
      } else {
        ShoppingCart.removeWhere((el) => el.id == index);
      }
      widget.updateCount();
    });
  }

  void increment(index) {
    setState(() {
      ShoppingCart.elementAt(ShoppingCart.indexWhere((el) => el.id == index))
          .count++;
      widget.updateCount();
    });
  }

  void decrement(index) {
    setState(() {
      if (ShoppingCart.elementAt(
              ShoppingCart.indexWhere((el) => el.id == index)).count >
          1) {
        ShoppingCart.elementAt(ShoppingCart.indexWhere((el) => el.id == index))
            .count--;
      } else {
        ShoppingCart.removeWhere((el) => el.id == index);
      }
      widget.updateCount();
    });
  }*/

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        backgroundColor: Colors.amber[200],
        appBar: AppBar(
          title: const Text('Избранное'),
          backgroundColor: Colors.white70,
        ),
        body: FutureBuilder<List<Items>>(
            future: ItemsFavList,
            builder: (context, snapshot) {
              if (snapshot.connectionState == ConnectionState.waiting) {
                return Center(child: CircularProgressIndicator());
              } else if (snapshot.hasError) {
                return Center(child: Text('Error: ${snapshot.error}'));
              } else if (!snapshot.hasData || snapshot.data!.isEmpty) {
                return Center(child: Text('No products found'));
              }

              final ItemsFavList = snapshot.data!;
              return ItemsFavList.length != 0
                  ? GridView.builder(
                      gridDelegate:
                          const SliverGridDelegateWithFixedCrossAxisCount(
                        crossAxisCount: 2,
                        childAspectRatio: 0.63,
                      ),
                      itemCount: ItemsFavList.length,
                      itemBuilder: (BuildContext context, int index) {
                        return GestureDetector(
                          onTap: () {
                            NavToItem(ItemsFavList.elementAt(index).id);
                          },
                          child: Padding(
                            padding: const EdgeInsets.only(
                                right: 5.0, left: 5.0, top: 2.0, bottom: 5.0),
                            child: Container(
                              height: MediaQuery.of(context).size.height * 0.25,
                              //height: MediaQuery.of(context).size.height * 0.47,
                              decoration: BoxDecoration(
                                color: const Color.fromARGB(255, 255, 246, 218),
                                borderRadius: BorderRadius.circular(7.0),
                              ),
                              child: Column(
                                children: [
                                  const SizedBox(
                                    height: 10,
                                  ),
                                  Center(
                                    child: Container(
                                      decoration: BoxDecoration(
                                        border: Border.all(
                                            color: Colors.grey, width: 1),
                                      ),
                                      child: Image.network(
                                        ItemsFavList.elementAt(index).image,
                                        width:
                                            MediaQuery.of(context).size.width *
                                                0.4,
                                        height:
                                            MediaQuery.of(context).size.width *
                                                0.4,
                                        fit: BoxFit.cover,
                                        loadingBuilder:
                                            (context, child, loadingProgress) {
                                          if (loadingProgress == null)
                                            return child;
                                          return const CircularProgressIndicator();
                                        },
                                        errorBuilder:
                                            (context, error, stackTrace) {
                                          return Container(
                                            width: MediaQuery.of(context)
                                                    .size
                                                    .width *
                                                0.4,
                                            height: MediaQuery.of(context)
                                                    .size
                                                    .width *
                                                0.4,
                                            color: Colors.amber[200],
                                            child: const Center(
                                                child: Text(
                                              'нет картинки',
                                              softWrap: true,
                                              textAlign: TextAlign.center,
                                            )),
                                          );
                                        },
                                      ),
                                    ),
                                  ),
                                  Padding(
                                    padding: const EdgeInsets.only(
                                        top: 4.0,
                                        bottom: 0.0,
                                        right: 15.0,
                                        left: 15.0),
                                    child: SizedBox(
                                      height: 35.0,
                                      child: Text(
                                        '${ItemsFavList.elementAt(index).name}',
                                        style: const TextStyle(fontSize: 12),
                                        softWrap: true,
                                        maxLines: 2,
                                        overflow: TextOverflow.ellipsis,
                                        textAlign: TextAlign.center,
                                      ),
                                    ),
                                  ),
                                  Padding(
                                    padding: const EdgeInsets.only(
                                        left: 15.0, right: 5.0),
                                    child: Row(children: [
                                      const Text(
                                        'Цена: ',
                                        style: TextStyle(fontSize: 12),
                                      ),
                                      Text(
                                        '${ItemsFavList.elementAt(index).cost} ₽',
                                        style: const TextStyle(
                                            fontSize: 12,
                                            color:
                                                Color.fromARGB(255, 6, 196, 9),
                                            fontWeight: FontWeight.bold),
                                      ),
                                      Expanded(
                                        child: Align(
                                          alignment: Alignment.centerRight,
                                          child: IconButton(
                                              onPressed: () => {
                                                    //AddFavorite(ItemsFavList.elementAt(index).id)
                                                  },
                                              icon: const Icon(Icons.favorite)),
                                        ),
                                      ),
                                    ]),
                                  ),
                                  ItemsFavList.elementAt(index).shopcart
                                      ? SizedBox(
                                          height: 40.0,
                                          width: MediaQuery.of(context)
                                                  .size
                                                  .width *
                                              0.4,
                                          child: Padding(
                                            padding: const EdgeInsets.only(
                                                left: 5.0, right: 5.0),
                                            child: Expanded(
                                              child: Row(
                                                  mainAxisAlignment:
                                                      MainAxisAlignment.center,
                                                  children: [
                                                    IconButton(
                                                        icon:
                                                            Icon(Icons.remove),
                                                        onPressed: () => {
                                                              //decrement(ItemsFavList.elementAt(index).id),
                                                            }),
                                                    Container(
                                                      height: 25.0,
                                                      width: 30.0,
                                                      decoration: BoxDecoration(
                                                        borderRadius:
                                                            BorderRadius
                                                                .circular(5.0),
                                                        border: Border.all(
                                                            color: const Color
                                                                .fromRGBO(
                                                                255, 160, 0, 1),
                                                            width: 2),
                                                      ),
                                                      child: Padding(
                                                        padding:
                                                            const EdgeInsets
                                                                .all(1.0),
                                                        child: Text(
                                                          ItemsFavList
                                                                  .elementAt(
                                                                      index)
                                                              .count
                                                              .toString(),
                                                          style:
                                                              const TextStyle(
                                                                  fontSize:
                                                                      14.0,
                                                                  color: Colors
                                                                      .black),
                                                          textAlign:
                                                              TextAlign.center,
                                                        ),
                                                      ),
                                                    ),
                                                    IconButton(
                                                        icon: Icon(Icons.add),
                                                        onPressed: () => {
                                                              //increment(ItemsFavList.elementAt(index).id),
                                                            }),
                                                  ]),
                                            ),
                                          ),
                                        )
                                      : Expanded(
                                          child: Align(
                                          alignment: Alignment.bottomCenter,
                                          child: Padding(
                                            padding: const EdgeInsets.only(
                                                bottom: 10.0),
                                            child: ElevatedButton(
                                              style: ElevatedButton.styleFrom(
                                                foregroundColor:
                                                    const Color.fromARGB(
                                                        255, 0, 0, 0),
                                                shape: RoundedRectangleBorder(
                                                    borderRadius:
                                                        BorderRadius.circular(
                                                            30),
                                                    side: const BorderSide(
                                                        width: 2,
                                                        color: Color.fromRGBO(
                                                            255, 160, 0, 1))),
                                                backgroundColor:
                                                    const Color.fromARGB(
                                                        255, 255, 246, 218),
                                              ),
                                              child: const Text("В корзину",
                                                  style:
                                                      TextStyle(fontSize: 12)),
                                              onPressed: () {
                                                //AddShopCart(ItemsFavList.elementAt(index).id);
                                              },
                                            ),
                                          ),
                                        ))
                                ],
                              ),
                            ),
                          ),
                        );
                      })
                  : const Center(child: Text('Нет избранных товаров'));
            }));
  }
}
