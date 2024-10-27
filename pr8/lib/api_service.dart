import 'dart:convert';

import 'package:dio/dio.dart';
import 'package:pr8/model/items.dart';

class ApiService {
  final Dio _dio = Dio();

  Future<List<Items>> getProducts() async {
    try {
      final response = await _dio.get('http://192.168.1.121:8080/products');
      if (response.statusCode == 200) {
        List<Items> products = (response.data as List)
            .map((product) => Items.fromJson(product))
            .toList();

        return products;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<List<Items>> getFavoriteProducts() async {
    try {
      final response =
          await _dio.get('http://192.168.1.121:8080/favorite_products');
      if (response.statusCode == 200) {
        List<Items> products = (response.data as List)
            .map((product) => Items.fromJson(product))
            .toList();

        return products;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<List<Items>> getShopCartProducts() async {
    try {
      final response =
          await _dio.get('http://192.168.1.121:8080/shop_cart_products');
      if (response.statusCode == 200) {
        List<Items> products = (response.data as List)
            .map((product) => Items.fromJson(product))
            .toList();

        return products;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<int> getCountShopCartProducts() async {
    try {
      final response =
          await _dio.get('http://192.168.1.121:8080/shop_cart_products');
      if (response.statusCode == 200) {
        int count = (response.data as List).toList().length;

        return count;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<Items> getProductsByID(int index) async {
    final link = 'http://192.168.1.121:8080/products/${index.toString()}';
    try {
      final response = await _dio.get(link);
      if (response.statusCode == 200) {
        Items product = Items.fromJson(response.data);
        print(product);
        return product;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<void> updateProductStatus(Items item) async {
    final link = 'http://192.168.1.121:8080/products/update/${item.id}';
    try {
      final response = await _dio.put(link, data: {
        'Name': item.name,
        'Image': item.image,
        'Cost': item.cost,
        'Describtion': item.describtion,
        'Favorite': item.favorite,
        'ShopCart': item.shopcart,
        'Count': item.count,
      });
      if (response.statusCode == 200) {
        return;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }

  Future<void> deleteProduct(int index) async {
    final link = 'http://192.168.1.121:8080/products/delete/${index}';
    try {
      final response = await _dio.delete(link);
      if (response.statusCode == 200) {
        return;
      } else {
        throw Exception('Failed to load products');
      }
    } catch (e) {
      throw Exception('Error fetching products: $e');
    }
  }
}
